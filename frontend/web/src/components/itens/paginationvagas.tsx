'use client'
import { VagaService } from "@/services/admin_service/vagas/vagas_service";
import { Button } from "../ui/button"
import { Pagination, PaginationContent, PaginationItem, PaginationPrevious, PaginationLink, PaginationEllipsis, PaginationNext } from "../ui/pagination"
import { useEffect, useState } from "react";
import { VagaData, VagasResponse } from "@/lib/types/vagatypes";

interface Props {
    onClickPrevius: () => void;
    onClickNext: () => void;
    Handle?: (page: number) => void;
    length: number;
    vagasPorPagina: number;
    currentPage: number;
}
export function PaginationVagas() {
    const [currentPage, setCurrentPage] = useState(1); // Página atual
    const [vagasPorPagina] = useState(5);
    const vagaService = new VagaService();
    const [vagas, setVagas] = useState<VagaData[]>([]);
    const [error, setError] = useState<string | null>(null);
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
    const paginatedVagas = () => {
        const startIndex = (currentPage - 1) * vagasPorPagina;
        const endIndex = startIndex + vagasPorPagina;
        return vagas.slice(startIndex, endIndex);
    };
    const handlePageChange = (page: number) => {
        if (page < 1) {
            return;
        }
        if (page > Math.ceil(vagas.length / vagasPorPagina)) {
            return;
        }
        setCurrentPage(page);
    };
    return (
        <div className="text-white">
            {error && <p>{error}</p>}

            <ul>
                {paginatedVagas().map((vaga) => (
                    <li key={vaga.numero_vaga}>
                        <strong>Número da Vaga:</strong> {vaga.numero_vaga}<br />
                        <strong>Tipo de Vaga:</strong> {vaga.tipo_vaga}<br />
                        <strong>Status:</strong> {vaga.status_vaga}<br />
                    </li>
                ))}
            </ul>
            <Pagination>
                <PaginationContent>
                    <PaginationItem>
                        <PaginationPrevious
                            onClick={() => handlePageChange(currentPage - 1)}
                        />
                    </PaginationItem>
                    {Array.from({ length: Math.ceil(vagas.length / vagasPorPagina) }, (_, index) => (
                        <PaginationItem key={index} className="text-black">
                            <PaginationLink
                                href="#"
                                onClick={() => handlePageChange(index + 1)}
                                isActive={currentPage === index + 1}
                            >
                                {index + 1}
                            </PaginationLink>
                        </PaginationItem>
                    ))}
                    {Math.ceil(vagas.length / vagasPorPagina) > 5 && (
                        <PaginationItem>
                            <PaginationEllipsis />
                        </PaginationItem>
                    )}

                    <PaginationItem>
                        <PaginationNext
                            onClick={() => handlePageChange(currentPage + 1)}
                            aria-disabled={currentPage === Math.ceil(vagas.length / vagasPorPagina)}
                        />
                    </PaginationItem>
                </PaginationContent>
            </Pagination>

            <Button onClick={() => vagaService.GetAllVaga()}>
                Puxar Vagas
            </Button>
        </div>
    )
}