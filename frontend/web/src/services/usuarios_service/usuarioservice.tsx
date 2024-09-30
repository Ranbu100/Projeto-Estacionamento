'use client'

import { axiosInstance } from "../axiosinstance";
import { AxiosResponse } from "axios";
import { LoginUserResponse, UserData } from "@/lib/types/usertypes";

interface LoginResp {
    message: string;
    token: string;
}

export class UsuarioService {
    async LoginUser(userData: UserData): Promise<LoginUserResponse> {
        try {
            const response = await axiosInstance.post<LoginResp, AxiosResponse<LoginResp, UserData>, UserData>("/api/v1/login", userData, {
                headers: {
                    'Content-Type': 'application/json'
                }

            });

            if (response.status == 200) {
                console.log('Login realizado com sucesso!');
                console.log(response.status);
                const { token } = response.data;
                console.log(token);
                return { success: true }
            } else {
                console.log('Erro no login:', response.status);
                throw new Error('Erro no login');
            }

        } catch (error) {
            console.error("Erro ao logar o usuario", error);
            return { success: false};
        }
    }
    GetAllUser() {
        return axiosInstance.get("/api/v1/usuarios")
    }
    async CreateUser(userData: any) {

        try {
            const response = await axiosInstance.post("/api/v1/usuarios", userData, {
                headers: {
                    'Content-Type': 'application/json'
                }
            })

            if (response.status == 201) {
                console.log('Cadastro realizado com sucesso!');
                console.log(response.status);
                return response.data
            } else {
                console.log('Erro no cadastro:', response.status);
                throw new Error('Erro no cadastro');
            }
        } catch (error) {
            console.error("Erro ao cadastrar o usuario", error);
        }

    }
}