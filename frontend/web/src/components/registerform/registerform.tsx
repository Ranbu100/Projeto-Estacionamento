'use client'
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Textinput } from "../itens/text_input"
import { Senhainput } from "../itens/input_senha"
import { useState } from "react"
import { Registerbutton } from "../itens/register_button"


export function Registerform(){
    const [isVisible, setIsVisible] = useState(false);
    return (
    <div className="flex justify-center items-center h-5/2 w-5/2">
        <Card className="w-[350px] bg-blue-900 rounded-lg" >
          <CardHeader>
            <CardTitle className="text-white font-extrabold text-3xl">Cadastro</CardTitle>
          </CardHeader>
          <CardContent>
            <form>
              <div className="grid w-full items-center gap-4">
                <Textinput inputmode="text" id="nome" labeltext="Nome" placeholder="Escreva o seu nome completo"/>
                <Textinput inputmode="email" id="email" labeltext="Email" placeholder="Escreva o seu email"/>
                <Senhainput label="Senha" isVisible={isVisible} onChangeVisibility={() => setIsVisible(!isVisible)}/>
                <Senhainput label="Confirme sua senha" isVisible={isVisible} onChangeVisibility={() => setIsVisible(!isVisible)}/>
                <Textinput inputmode="tel" id="telefone" labeltext="Telefone" placeholder="Somente números e DDD" />  
              </div>
            </form>
          </CardContent>
          <CardFooter className="flex justify-end">
          <Registerbutton link="/" classname="text-blue-900 bg-white"/>
          </CardFooter>
        </Card>
    </div>
    )
}