# Etap budowania
FROM node:16-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Etap uruchomieniowy - użycie obrazu serwera statycznego np. serve
FROM node:16-alpine
RUN npm install -g serve
WORKDIR /app
COPY --from=build /app/build ./build
EXPOSE 3000
CMD ["serve", "-s", "build", "-l", "3000"]
