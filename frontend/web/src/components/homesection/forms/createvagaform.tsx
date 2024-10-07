'use client'
import { RenderField } from "@/components/itens/forms_field";
import {
    Card,
    CardContent,
    CardDescription,
    CardTitle,
    CardHeader,
} from "@/components/ui/card";
import { yupResolver } from "@hookform/resolvers/yup";
import { formSchemaCreate } from "./validator/createvaga_validator";
import { useForm } from "react-hook-form";
import { VagaService } from "@/services/admin_service/vagas/vagas_service";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import Dropdown from "@/components/itens/dropdown";

export function CriarVagasForm() {
    const { control, handleSubmit, formState: { errors } } = useForm({
        resolver: yupResolver(formSchemaCreate),
      });
      const [isCreate, setIsCreate] = useState(false);
      const tipos = [
        { label: 'Carro', value: 1 },
        { label: 'Moto', value: 2 }
      ];
      const onSubmit = async (data: { tipo_vaga: string; numero_vaga: number}) => {
        try {
            const createdvaga = await new VagaService().CreateVaga({
                numero_vaga: data.numero_vaga,
                tipo_vaga: (data.tipo_vaga == '1' ? 1 : 2),
                status_vaga: 1,
            },);
            setIsCreate(createdvaga.success);
        } catch (error) {
            alert("A vaga não foi criada");
            console.error(error);
        }
      }
    return (
        <div className="w-vw h-dvh p-12 bg-sky-700 bg-gradient-to-tr bg-right bg-no-repeat bg-fixed">
            <Card className="h-1/2 w-1/2 bg-blue-900">
                <CardHeader className="text-white">
                    <CardTitle>Crie sua Vaga</CardTitle>
                    <CardDescription className="text-white">Insira os dados corretos referentes à vaga a ser criada</CardDescription>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmit(onSubmit)} className="space-y-2">
                    <RenderField control={control} name="numero_vaga" label="Numero da Vaga" errors={errors} type={undefined} /> 
                    <Dropdown label="Selecione o tipo da vaga" options={tipos} name="tipo_vaga" control={control}></Dropdown>
                    <Button className="bg-white text-blue-900 " type="submit" >Criar</Button>
                    </form>
                   
                </CardContent>
            </Card>
        </div>

    )
}