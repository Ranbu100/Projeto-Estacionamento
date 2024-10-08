interface VagaData {
    numero_vaga: number;
    tipo_vaga: number;
    status_vaga: number;
}
interface VagasResponse {
    success: boolean;
    message?: string;
    vaga? : VagaData;
    data?: VagaData[];
}
export type {
    VagaData,
    VagasResponse
}