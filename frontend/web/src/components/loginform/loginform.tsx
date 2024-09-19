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
import { Loginbutton } from "../itens/loginbutton"
import { useState } from "react"
export function Loginform() {
  const [isVisible, setIsVisible] = useState(false);
  return (
    <div className="flex justify-center items-center h-5/2 w-5/2">
        <Card className="w-[350px] bg-blue-900 rounded-lg" >
          <CardHeader>
            <CardTitle className="text-white font-extrabold text-3xl">Logar</CardTitle>
          </CardHeader>
          <CardContent>
            <form>
              <div className="grid w-full items-center gap-4">
                <Textinput inputmode="email" id="email" labeltext="Email" placeholder="Escreva o seu email"/>
                <Senhainput label="Senha" isVisible={isVisible} onChangeVisibility={() => setIsVisible(!isVisible)}/>
              </div>
            </form>
          </CardContent>
          <CardFooter className="flex justify-end">
            <Loginbutton classname="text-blue-900" link="/"/>
          </CardFooter>
        </Card>
    </div>
    
  );
}