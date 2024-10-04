interface VagaData {
    numero_vaga: number;
    tipo_vaga: string;
    status_vaga: number;
}
interface VagasResponse {
    success: boolean;
    message?: string;
}
export type {
    VagaData,
    VagasResponse
}