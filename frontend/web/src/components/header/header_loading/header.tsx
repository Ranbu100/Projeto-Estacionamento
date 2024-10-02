"use client"
import Image from "next/image";
import Logo_header from "../../../../public/images/logo_header.svg";
import { ItemMenu } from "./itemMenu";
import { Registerbutton } from "./register_button";
import { useRouter } from "next/navigation";
import { Loginbutton } from "./loginbutton";
export function Header() {
    const  Router  = useRouter();
    return (
        <header className="flex items-center w-full h-24 bg-blue-900">
            <div className="flex flex-wrap items-center w-full max-w-[1246px] px-4 mx-auto justify-between">
                <div className="flex items-center gap-2 sm:gap-4">
                    <Image src={Logo_header} alt="Logo" className="w-auto h-12 sm:w-auto sm:h-auto" />
                    <div className="hidden md:flex md:items-center md:gap-6">
                        <button className="md:text-white  md:text-2xl md:font-bold"
                        onClick={() => Router.push("/")}
                        >Estacionamentos LTDA
                        </button>
                    </div>
                </div>

                <div className="flex items-center gap-2">
                    <ul className="flex items-center gap-2 sm:gap-1 justify-between">
                        <ItemMenu nome="Sobre nós" route="/login" />
                        <ItemMenu nome="Contato" route="/login"/>
                        <ItemMenu nome="Planos" route="/login"/>
                    </ul>
                </div>
                <div className="flex items-center gap-1 sm:gap-3">
                    <Loginbutton 
                    classname="flex items-center font-semibold text-blue-900 h-auto w-auto p-0 lg:p-2 bg-white border rounded-sm" 
                    classname_text="h-auto w-auto p-0.5" 
                    link="/login"
                    />
                    <Registerbutton 
                    classname="flex items-center font-semibold text-white bg-transparent border border-white h-auto w-auto  p-0 lg:p-2 rounded-sm" 
                    classname_text="h-auto w-auto p-0.5" 
                    link="/register"
                    />
                </div>
            </div>
        </header>
    );
}
