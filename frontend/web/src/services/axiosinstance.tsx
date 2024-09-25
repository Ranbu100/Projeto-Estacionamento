import axios from "axios";


export const axiosInstance = axios.create({
    baseURL: "http://nicholas:8080",
})