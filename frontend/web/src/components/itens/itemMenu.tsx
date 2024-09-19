'use client'
import { useRouter } from "next/navigation";
import { Button } from "../ui/button";

type Props = {
    nome: string;
    route: string;
}

export function ItemMenu({nome, route} : Props){
    const Router = useRouter();
    return (
        <Button variant="link"
        className="flex items-center gap-1 sm:gap-2"
        onClick={() => Router.push(route)}>
            <span className="text-white text-xs">{nome}</span>
        </Button>
    )
}