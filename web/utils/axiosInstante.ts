import axios from "axios";
import { BaseURL } from "./baseURL";

export const axiosInstante = axios.create({
    baseURL: BaseURL,
})