'use client'
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

import { RenderField } from "../itens/forms_field";
import { useForm } from "react-hook-form";
import { formSchema } from "./login_validator";
import { yupResolver } from "@hookform/resolvers/yup"
import { UsuarioService } from "@/services/usuarios_service/usuarioservice";
import { useState } from "react";
import { Button } from "../ui/button";
import { AlertLoginConfirm } from "./alertlogin";
import { AlertLoginErrors } from "./errors";
export function Loginform() {
  const { control, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(formSchema),
  });
  const [isLogged, setIsLogged] = useState(false);
  const [loginError, setLoginError] = useState(false);
  const onSubmit = async (data: {email: string, senha: string}) => {
    try {
      const logged = await new UsuarioService().LoginUser({
        email: data.email,
        senha: data.senha,
      });

      setIsLogged(logged.success);
      setLoginError(!logged.success);

    } catch (error) {
      console.error('Erro ao logar usuário:', error);
      alert('Erro ao realizar login. Tente novamente.');
    }
  };
  return (
    <div className="flex justify-center items-center h-4/5 w-screen">
        <Card className="w-[350px] bg-blue-900 rounded-lg" >
          <CardHeader>
            <CardTitle className="text-white font-extrabold text-3xl">Logar</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-2">
              <RenderField control={control} name="email" label="Email" errors={errors} type={undefined} />
              <RenderField control={control} name="senha" label="Senha" errors={errors} type={"senha"} />              
              <Button className="bg-white text-blue-900 " type="submit">Logar</Button>
            </form>
            {isLogged ? (
            <AlertLoginConfirm show={isLogged} onClose={() => setIsLogged(false)} />
          ) : loginError ? (
            <AlertLoginErrors show={true} onClose={() => setLoginError(false)} />
          ) : null}

          </CardContent>
        </Card>
    </div>
    
  );
}