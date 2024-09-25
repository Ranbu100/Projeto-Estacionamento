'use client';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Textinput } from "../itens/text_input";
import { Senhainput } from "../itens/input_senha";
import { useState } from "react";
import { Registerbutton } from "../itens/register_button";
import { UsuarioService } from "@/services/usuarios_service/usuarioservice";

export function Registerform() {
  const [isVisible, setIsVisible] = useState(false);
  
  // Estado para armazenar os valores do formulário
  const [formData, setFormData] = useState({
    nome: '',
    email: '',
    senha: '',
    confirmaSenha: '',
    telefone: ''
  });

  // Função para atualizar os campos do formulário
  const handleInputChange = (e: any) => {
    const { id, value } = e.target;
    setFormData({
      ...formData,
      [id]: value // Atualiza o campo correspondente no estado
    });
  };

  // Função para manipular o envio do formulário
  const handleSubmit = async (e: any) => {
    e.preventDefault(); // Previne o comportamento padrão de recarregar a página

    // Validação de senha
    if (formData.senha !== formData.confirmaSenha) {
      alert('As senhas não conferem!'); 
      return; 
    }

    try {
      await new UsuarioService().CreateUser({
        nome: formData.nome,
        email: formData.email,
        senha: formData.senha,
        telefone: formData.telefone
      });
      alert('Cadastro realizado com sucesso!');
    } catch (error) {
      console.error('Erro ao cadastrar usuário:', error);
      alert('Erro ao realizar cadastro. Tente novamente.');
    }
  };

  return (
    <div className="flex justify-center items-center h-5/2 w-5/2">
      <Card className="w-[350px] bg-blue-900 rounded-lg">
        <CardHeader>
          <CardTitle className="text-white font-extrabold text-3xl">Cadastro</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <div className="grid w-full items-center gap-4">
              <Textinput
                inputmode="text"
                id="nome"
                labeltext="Nome"
                placeholder="Escreva o seu nome completo"
                value={formData.nome}
                onChange={handleInputChange}
              />
              <Textinput
                inputmode="email"
                id="email"
                labeltext="Email"
                placeholder="Escreva o seu email"
                value={formData.email}
                onChange={handleInputChange}
              />
              <Senhainput
                label="Senha"
                isVisible={isVisible}
                onChangeVisibility={() => setIsVisible(!isVisible)}
                value={formData.senha}
                onChange={(e) => setFormData({ ...formData, senha: e.target.value })}
              />
              <Senhainput
                label="Confirme sua senha"
                isVisible={isVisible}
                onChangeVisibility={() => setIsVisible(!isVisible)}
                value={formData.confirmaSenha}
                onChange={(e) => setFormData({ ...formData, confirmaSenha: e.target.value })}
              />
              <Textinput
                inputmode="tel"
                id="telefone"
                labeltext="Telefone"
                placeholder="Somente números e DDD"
                value={formData.telefone}
                onChange={handleInputChange}
              />
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex justify-end">
          <Registerbutton onClick={handleSubmit} type_button="submit" link="/" classname="text-blue-900 bg-white" />
        </CardFooter>
      </Card>
    </div>
  );
}
