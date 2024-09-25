
import { axiosInstance } from "../axiosinstance";



export class UsuarioService {
    GetAllUser() {
        return axiosInstance.get("/api/v1/usuarios")
    }
    CreateUser(userData: any) {
        return axiosInstance.post("/api/v1/usuarios", userData, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.status === 200) {
                console.log('Cadastro realizado com sucesso!');
                return response.data; // Retorna a resposta se necessário
            } else {
                console.log('Erro no cadastro:', response.status);
                throw new Error('Erro no cadastro');
            }
        }).catch(function (error){
            console.log(error);
        })
    }
}