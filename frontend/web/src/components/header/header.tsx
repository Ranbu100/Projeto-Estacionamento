import Image from "next/image";
import Logo_header from "./../../assets/images/logo_header.svg";
import { ItemMenu } from "./itemMenu";
import { Login } from "./login_button";
import { Cadastro } from "./register_button";

export function Header() {
    return (
        <header className="flex items-center w-full h-24 bg-blue-900">
            <div className="flex flex-wrap items-center w-full max-w-[1246px] px-4 mx-auto justify-between">
                <div className="flex items-center gap-4 sm:gap-14">
                    <Image src={Logo_header} alt="Logo" className="w-10 h-10 sm:w-auto sm:h-auto" />
                    <h1 className="text-white text-2xl md:text-3xl font-bold ">Estacionamentos LTDA</h1>
                </div>

                <div className="hidden lg:flex items-center gap-8">
                    <ul className="flex items-center gap-4 sm:gap-8 justify-between">
                        <ItemMenu nome="Sobre nós" />
                        <ItemMenu nome="Contato" />
                        <ItemMenu nome="Planos" />
                    </ul>
                </div>
                <div className="flex items-center gap-3 sm:gap-5">
                    <Login/>
                    <Cadastro/>
                </div>
            </div>
        </header>
    );
}
