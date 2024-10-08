'use client'
import { Button } from "../ui/button";
import { useRouter } from "next/navigation";

export function HomeSection() {
    const router = useRouter();
    return (
        <div className="w-vw h-dvh p-12 bg-sky-700 bg-gradient-to-tr bg-right bg-no-repeat bg-fixed">
            <div className="bg-black bg-opacity-30 flex flex-col w-1/2 h-1/2 p-2 border border-spacing-3 border-black rounded-2xl">
                <span className="text-xl sm:text-3xl text-white font-extrabold justify-normal">Crie, Gerencie e Monitore todo o seu Estacionamento</span>
                <div className="flex flex-wrap text-white">
                    <Button onClick={() => { router.push("/home/criarvagas") }} variant={"ghost"} className="h-1/6 w-1/4 p-10">
                        <span>Crie vagas</span>
                    </Button>
                    <Button variant={"ghost"} onClick={() => { router.push("/home/gerenciarvagas") }} className="h-1/6 w-1/4 p-10">
                        <span>Gerencie vagas</span>
                    </Button>
                    <Button variant={"ghost"} className="h-1/6 w-1/4 p-10">
                        <span>Monitore vagas</span>
                    </Button>
                </div>

            </div>

        </div>
    )
}