"use client"
import Image from "next/image";
import Logo_header from "../../../../public/images/logo_header.svg";

import { useRouter } from "next/navigation";
import { HomeDrawer } from "./home_drawer";

export function Header_home() {
    const Router = useRouter();
    return (
        <header className="flex items-center w-vw h-24 bg-blue-900">
            <div className="flex flex-wrap items-center w-full max-w-[1246px] px-4 mx-auto bg-no-repeat justify-between">
                <div className="flex items-center gap-2 sm:gap-4">
                    <Image src={Logo_header} alt="Logo" className="w-auto h-12 sm:w-auto sm:h-auto" />
                </div>
                <div className="flex items-center gap-2">

                </div>
                <div className="flex items-center gap-1 sm:gap-3 w-auto h-auto">
                        
                        <div className="justify-end">
                        <HomeDrawer />
                        </div>
                        

                </div>
            </div>
        </header>
    );
}
