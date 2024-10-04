'use client'
import { VagaData, VagasResponse } from "@/lib/types/vagatypes";
import { axiosInstance } from "@/services/axiosinstance";
import { AxiosResponse } from "axios";
import nookies from 'nookies';
interface VagaResp {
    message: string;
}
export class VagaService{
    async CreateVaga(vagaData: VagaData, ctx?: any): Promise<VagasResponse>{
        
       try {
        const cookies = nookies.get(ctx);
        const token = cookies.token;
         const response = await axiosInstance.post("api/v1/vagas",vagaData,
            { headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json'
            }}
         )
         
         if (response.status == 200) {
            console.log('Vaga criada com sucesso!');
            console.log(response.status);
            return { success: true }
        } else {
            console.log('Erro:', response.status);
            throw new Error('Erro ao criar a vaga');
        }
       } catch (error) {
        console.log(vagaData);
        console.error("Erro ao Criar a classe", error);
            return { success: false}
       }
    }
}