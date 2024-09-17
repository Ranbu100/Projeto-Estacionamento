'use client'
import { useRouter } from "next/navigation";

type Props = {
    nome: string;
    route: string;
}

export function ItemMenu({nome} : Props, {route}: Props){
    const Router = useRouter();
    return (
        <button 
        type="button"
        className="flex items-center gap-1 sm:gap-3"
        onClick={() => Router.push(route)}>
            <span className="text-white">{nome}</span>
        </button>
    )
}