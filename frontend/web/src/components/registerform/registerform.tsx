"use client"

import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Button } from "@/components/ui/button";
import { formSchema } from "./register_validator";
import { RenderField } from "../itens/forms_field";
import { UsuarioService } from "@/services/usuarios_service/usuarioservice";
import { AlertRegisterConfirm } from "./register_confirm";
import { Card, CardHeader, CardTitle, CardContent } from "../ui/card";
import React, { useState } from 'react';
import { formatPhoneNumber, unformatPhoneNumber } from '../../lib/mask/maskphone'; // importe as funções
import { AlertRegisterErrors } from "./errors";

export default function RegisterForm() {

  const { control, handleSubmit, setValue, watch, formState: { errors } } = useForm({
    resolver: yupResolver(formSchema),
    defaultValues: {
      nome: '',
      email: '',
      senha: '',
      confirmaSenha: '',
      telefone: ''
    },
  });

  const [isRegistered, setIsRegistered] = useState(false);
  const [registrationError, setRegistrationError] = useState(false);  // Novo estado para o erro

  // Watch para observar o campo telefone e aplicar a máscara automaticamente
  const phoneValue = watch('telefone');

  // Função para formatar o valor ao digitar
  const handlePhoneChange = (value: string) => {
    const formattedPhone = formatPhoneNumber(value);
    setValue('telefone', formattedPhone); // Atualiza o valor com a máscara aplicada
  };

  const onSubmit = async (data: { nome: string, email: string, senha: string, telefone: string }) => {
    try {
      // Desformata o número de telefone antes de enviar
      const phoneToSubmit = unformatPhoneNumber(data.telefone);

      const registred = await new UsuarioService().CreateUser({
        nome: data.nome,
        email: data.email,
        senha: data.senha,
        telefone: phoneToSubmit,
        tipo_usuario: "Comum"
      });

      if (registred) {
        setIsRegistered(true);
        setRegistrationError(false);  // Resetar erro
      } else {
        setIsRegistered(false);
        setRegistrationError(true);  // Mostrar erro
      }
    } catch (error) {
      console.error('Erro ao cadastrar usuário:', error);
      setIsRegistered(false);
      setRegistrationError(true);  // Exibir erro em caso de falha
    }
  };

  return (
    <div className="flex justify-center items-center h-5/3 w-screen">
      <Card className="w-[350px] bg-blue-900 rounded-lg">
        <CardHeader>
          <CardTitle className="text-white font-extrabold text-3xl">Cadastro</CardTitle>
        </CardHeader>
        <CardContent>

          <form onSubmit={handleSubmit(onSubmit)} className="space-y-2">
            <RenderField control={control} name="nome" label="Nome" errors={errors} type={undefined} />
            <RenderField control={control} name="email" label="Email" errors={errors} type={undefined} />
            <RenderField control={control} name="senha" label="Senha" type="senha" errors={errors} />
            <RenderField control={control} name="confirmaSenha" label="Confirme sua senha" type="confirmarsenha" errors={errors} />
            {/* Aplica o formato ao campo de telefone */}
            <RenderField
              control={control}
              name="telefone"
              label="Telefone"
              errors={errors}
              type={'telefone'}
              onChange={(e: any) => handlePhoneChange(e.target.value)}
              value={phoneValue}
              maxLength={15}
            />
            <Button className="bg-white text-blue-900 py-1" type="submit">Registrar</Button>
          </form>

          {isRegistered ? (
            <AlertRegisterConfirm show={isRegistered} onClose={() => setIsRegistered(false)} />
          ) : registrationError ? (
            <AlertRegisterErrors show={true} onClose={() => setRegistrationError(false)} />
          ) : null}

        </CardContent>
      </Card>
    </div>
  );
}
