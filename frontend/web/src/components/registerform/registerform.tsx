
"use client"

import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { formSchema } from "./register_validator";
import { RenderField } from "../itens/forms_field";
import { UsuarioService } from "@/services/usuarios_service/usuarioservice";
import { AlertRegisterConfirm } from "./register_confirm";
import { Card, CardHeader, CardTitle, CardContent } from "../ui/card";

export default function RegisterForm() {
  const { control, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(formSchema),
    defaultValues: {
      nome: '',
      email: '',
      senha: '',
      confirmaSenha: ''
    },
  });

  const [isRegistered, setIsRegistered] = useState(false);

  const onSubmit = async (data: { nome: string, email: string, senha: string, telefone: string }) => {
    try {
      const registred = await new UsuarioService().CreateUser({
        nome: data.nome,
        email: data.email,
        senha: data.senha,
        telefone: data.telefone,
        tipo_usuario: "Comum"
      });

      if (registred) {
        setIsRegistered(true);
      }
    } catch (error) {
      console.error('Erro ao cadastrar usuário:', error);
      alert('Erro ao realizar cadastro. Tente novamente.');
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
        <RenderField control={control} name="nome" label="Nome" errors={errors} type={undefined}  />
        <RenderField control={control} name="email" label="Email" errors={errors} type={undefined}  />
        <RenderField control={control} name="senha" label="Senha" type="senha" errors={errors}  />
        <RenderField control={control} name="confirmaSenha" label="Confirme sua senha" type="confirmarsenha" errors={errors} />
        <RenderField control={control} name="telefone" label="Telefone" errors={errors} type={'telefone'}  />
        <Button className="bg-white text-blue-900 py-1" type="submit">Registrar</Button>
      </form>

      {isRegistered && (
        <AlertRegisterConfirm show={isRegistered} onClose={() => setIsRegistered(false)} />
      )}
      </CardContent>
    </Card>
  </div>
  );
}