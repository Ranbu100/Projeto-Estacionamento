'use client'
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Textinput } from "./text_input"
import { Senhainput } from "./input_senha"
import { Loginbutton } from "./loginbutton"
import { useState } from "react"
export function Loginform() {
  const [isVisible, setIsVisible] = useState(false);
  return (
    <div className="flex justify-center items-center h-5/2 w-5/2">
        <Card className="w-[350px] bg-blue-900 rounded-lg" >
          <CardHeader>
            <CardTitle className="text-white font-extrabold text-3xl">Login</CardTitle>
          </CardHeader>
          <CardContent>
            <form>
              <div className="grid w-full items-center gap-4">
                <Textinput id="email" labeltext="Email" placeholder="Escreva o seu email"/>
                <Senhainput isVisible={isVisible} onChangeVisibility={() => setIsVisible(!isVisible)}/>
              </div>
            </form>
          </CardContent>
          <CardFooter className="flex justify-end">
            <Loginbutton/>
          </CardFooter>
        </Card>
    </div>
    
  );
}