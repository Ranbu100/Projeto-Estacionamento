'use client'
import { VagaData, VagasResponse } from "@/lib/types/vagatypes";
import { axiosInstance } from "@/services/axiosinstance";
import nookies from 'nookies';

export class VagaService{
    async CreateVaga(vagaData: VagaData): Promise<VagasResponse>{
        
       try {
        const cookies = nookies.get();
        const token = cookies.token;
         const response = await axiosInstance.post("api/v1/vagas",vagaData,
            { headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json'
            }}
         )
         console.log(vagaData);
         
         if (response.status == 201) {
            console.log('Vaga criada com sucesso!');
            console.log(response.status);
            return { success: true }
        } else {
            console.log('Erro:', response.status);
            throw new Error('Erro ao criar a vaga');
        }
       } catch (error) {
        console.log(vagaData);
        console.error("Erro ao Criar a Vaga", error);
            return { success: false}
       }
    }
    async GetAllVaga(): Promise<VagasResponse> {
        try {
            const cookies = nookies.get();
            const token = cookies.token;
            const response = await axiosInstance.get("api/v1/vagas", {
                headers: {
                    Authorization: `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
            });
    
            if (response.status === 200) {
                const vagasData = response.data;
    
                if (vagasData && typeof vagasData === 'object') {
                    // Converte o objeto em um array de VagaData
                    const allVagas: VagaData[] = Object.values(vagasData);
                    return { success: true, data: allVagas };
                } else {
                    return { success: false, message: "Dados mal formatados." };
                }
            } else {
                return { success: false, message: "Erro ao carregar vagas" };
            }
        } catch (error) {
            console.error("Erro ao Puxar a Vaga", error);
            return { success: false, message: "Erro ao carregar vagas" };
        }
    }
    
    
}