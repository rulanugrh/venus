import { axiosInstante } from "./axiosInstante";

export const VerifyToken = async (token: string) => {
    try{
        const response = await axiosInstante.post("/validate-token", {
            token,
        })

        return response.data;
    } catch (e: any) {
        console.log(e.message)
    }
};