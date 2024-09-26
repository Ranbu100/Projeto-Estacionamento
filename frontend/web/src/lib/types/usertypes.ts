interface UserData {
    nome?: string,
    email: string,
    senha: string,
    telefone?: string,
    tipo_usuario?: string,
}

interface LoginUserResponse {
    success: boolean;
    message?: string;
}

export type {
    UserData,
    LoginUserResponse,
}