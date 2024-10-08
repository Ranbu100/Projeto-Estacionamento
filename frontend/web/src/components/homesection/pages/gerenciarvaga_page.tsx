'use client'
import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card";
import { VagaData, VagasResponse } from "@/lib/types/vagatypes";
import { VagaService } from "@/services/admin_service/vagas/vagas_service";
import { PaginationVagas } from "@/components/itens/paginationvagas";

const vagaService = new VagaService();

export function GerenciarVaga() {
    const [vagas, setVagas] = useState<VagaData[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [currentPage, setCurrentPage] = useState(1); // Página atual
    const [vagasPorPagina] = useState(5); // Quantidade de vagas por página

    useEffect(() => {
        const fetchVagas = async () => {
            const response: VagasResponse = await vagaService.GetAllVaga();

            if (response.success && response.data) {
                setVagas(response.data);
            } else {
                setError(response.message || "Erro ao carregar as vagas.");
            }
        };

        fetchVagas();
    }, []);

    // Função para calcular quais vagas devem ser exibidas de acordo com a página atual
    const paginatedVagas = () => {
        const startIndex = (currentPage - 1) * vagasPorPagina;
        const endIndex = startIndex + vagasPorPagina;
        return vagas.slice(startIndex, endIndex);
    };

    return (
        <div className="w-vw h-dvh p-12 bg-sky-700 bg-gradient-to-tr bg-right bg-no-repeat bg-fixed flex justify-center items-center">
            <Card className="h-5/6 w-5/6 bg-blue-900 bg-no-repeat bg-fixed">
                <CardHeader className="text-white">
                    <CardTitle>Suas Vagas</CardTitle>
                    <CardDescription className="text-white">Gerencie suas vagas</CardDescription>
                </CardHeader>
                <CardContent>
                    <PaginationVagas/>
                </CardContent>
            </Card>
        </div>
    );
}
