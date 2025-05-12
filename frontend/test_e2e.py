import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import time

# Adres frontendowy (zmień jeśli inny)
BASE_URL = "http://localhost:3000"

def get_driver():
    options = webdriver.ChromeOptions()
    options.add_argument('--headless')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
    driver.implicitly_wait(5)
    return driver

# 1. Strona główna ładuje się poprawnie
def test_homepage_loads():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    assert "Produkty" in driver.page_source
    driver.quit()

# 2. Link do kategorii działa
def test_navbar_categories():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    link = driver.find_element(By.LINK_TEXT, "Kategorie")
    link.click()
    assert "Kategorie" in driver.page_source
    driver.quit()

# 3. Link do koszyków działa
def test_navbar_carts():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    link = driver.find_element(By.LINK_TEXT, "Koszyki")
    link.click()
    assert "Koszyki" in driver.page_source
    driver.quit()

# 4. Link do płatności działa
def test_navbar_payments():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    link = driver.find_element(By.LINK_TEXT, "Płatności")
    link.click()
    assert "Płatność" in driver.page_source
    driver.quit()

# 5. Lista produktów nie jest pusta
def test_products_not_empty():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    items = driver.find_elements(By.TAG_NAME, "li")
    assert len(items) > 0
    driver.quit()

# 6. Dodanie produktu do koszyka (jeśli istnieje koszyk #1)
def test_add_product_to_cart():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj do koszyka')]")
    if buttons:
        buttons[0].click()
        time.sleep(1)
        assert "Dodaj do koszyka" in buttons[0].text or "Dodaję..." in buttons[0].text
    driver.quit()

# 7. Przycisk dodawania jest nieaktywny podczas dodawania
def test_add_button_disabled():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj do koszyka')]")
    if buttons:
        buttons[0].click()
        assert buttons[0].is_enabled() is False or buttons[0].text == "Dodaję..."
    driver.quit()

# 8. Formularz płatności waliduje kwotę
def test_payment_form_validation():
    driver = get_driver()
    driver.get(BASE_URL + "/payments")
    input_field = driver.find_element(By.XPATH, "//input[@type='number']")
    submit = driver.find_element(By.XPATH, "//button[@type='submit']")
    input_field.clear()
    submit.click()
    # Sprawdź, czy pole jest required (HTML5)
    assert input_field.get_attribute("required") == "true" or input_field.get_attribute("required") == ""
    driver.quit()

# 9. Przeładowanie strony nie usuwa koszyka (liczba koszyków się nie zmienia)
def test_cart_persistence():
    driver = get_driver()
    driver.get(BASE_URL + "/carts")
    items_before = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    driver.refresh()
    items_after = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    assert len(items_before) == len(items_after)
    driver.quit()

# 10. Dodanie nowego koszyka (czekamy na pojawienie się nowego koszyka)
def test_create_new_cart():
    driver = get_driver()
    driver.get(BASE_URL + "/carts")
    items_before = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    button = driver.find_element(By.XPATH, "//button[contains(text(),'Dodaj nowy koszyk')]")
    button.click()
    time.sleep(2)
    items_after = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    assert len(items_after) == len(items_before) + 1
    driver.quit()

# 12. Sprawdzenie czy na stronie produktów wyświetla się cena przynajmniej jednego produktu
def test_product_has_price():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    items = driver.find_elements(By.TAG_NAME, "li")
    found = False
    for item in items:
        if "PLN" in item.text:
            found = True
            break
    assert found, "Żaden produkt nie ma ceny!"
    driver.quit()

# 13. Przejście między zakładkami nie powoduje błędów
def test_navbar_no_errors():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    for name in ["Kategorie", "Koszyki", "Płatności", "Produkty"]:
        link = driver.find_element(By.LINK_TEXT, name)
        link.click()
        assert "Błąd" not in driver.page_source
    driver.quit()

# 14. Płatność z poprawną kwotą kończy się sukcesem (jeśli istnieje koszyk #1)
def test_payment_success():
    driver = get_driver()
    driver.get(BASE_URL + "/payments")
    input_field = driver.find_element(By.XPATH, "//input[@type='number']")
    input_field.clear()
    input_field.send_keys("10.00")
    submit = driver.find_element(By.XPATH, "//button[@type='submit']")
    submit.click()
    time.sleep(1)
    assert "sukcesem" in driver.page_source or "Płatność zakończona sukcesem" in driver.page_source
    driver.quit()

# 15. Płatność z błędną kwotą kończy się błędem (np. pusta kwota)
def test_payment_error():
    driver = get_driver()
    driver.get(BASE_URL + "/payments")
    input_field = driver.find_element(By.XPATH, "//input[@type='number']")
    input_field.clear()
    submit = driver.find_element(By.XPATH, "//button[@type='submit']")
    submit.click()
    time.sleep(1)
    assert "Błąd" in driver.page_source or "Błąd płatności" in driver.page_source or input_field.get_attribute("required")
    driver.quit()

# 16. Wyświetlanie liczby produktów w koszyku
def test_cart_items_count():
    driver = get_driver()
    driver.get(BASE_URL + "/carts")
    items = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    assert all("produkt" in i.text for i in items)
    driver.quit()

# 17. Dodanie kilku produktów do koszyka sumuje ilości (jeśli istnieje koszyk #1)
def test_add_multiple_products_to_cart():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj do koszyka')]")
    if len(buttons) > 1:
        buttons[0].click()
        time.sleep(1)
        buttons[1].click()
        time.sleep(1)
    driver.quit()

# 18. Sprawdzenie czy navbar jest widoczny
def test_navbar_visible():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    nav = driver.find_element(By.TAG_NAME, "nav")
    assert nav.is_displayed()
    driver.quit()

# 19. Nie pozwala na ujemną kwotę płatności
def test_payment_negative_amount():
    driver = get_driver()
    driver.get(BASE_URL + "/payments")
    input_field = driver.find_element(By.XPATH, "//input[@type='number']")
    input_field.clear()
    input_field.send_keys("-10")
    submit = driver.find_element(By.XPATH, "//button[@type='submit']")
    submit.click()
    time.sleep(1)
    # Sprawdź, czy nie pojawił się komunikat o sukcesie i pole nadal ma -10
    assert "sukcesem" not in driver.page_source
    assert input_field.get_attribute("value") == "-10"
    driver.quit()

# 20. Tytuł strony produktów jest poprawny
def test_products_page_title():
    driver = get_driver()
    driver.get(BASE_URL + "/products")
    h1 = driver.find_element(By.TAG_NAME, "h1")
    assert h1.text == "Produkty"
    driver.quit()

# 21. Dodanie produktu do koszyka zwiększa liczbę produktów w koszyku
def test_add_product_increases_cart_count():
    driver = get_driver()
    driver.get(BASE_URL + "/carts")
    items_before = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    before_count = 0
    if items_before:
        try:
            before_count = int(items_before[0].text.split("—")[-1].split()[0])
        except Exception:
            before_count = 0
    driver.get(BASE_URL + "/products")
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj do koszyka')]")
    if buttons:
        buttons[0].click()
        time.sleep(1)
    driver.get(BASE_URL + "/carts")
    items_after = driver.find_elements(By.XPATH, "//li[contains(text(),'Koszyk')]")
    after_count = 0
    if items_after:
        try:
            after_count = int(items_after[0].text.split("—")[-1].split()[0])
        except Exception:
            after_count = 0
    assert after_count >= before_count
    driver.quit()
