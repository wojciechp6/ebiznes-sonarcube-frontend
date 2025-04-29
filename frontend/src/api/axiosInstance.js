// src/api/axiosInstance.js
import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: 'http://localhost:8080/api', // Upewnij się, że port i endpoint odpowiadają konfiguracji backendu
    headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*'
    }
});

export default axiosInstance;
