import Link from "next/link";
import { Button } from "../ui/button"

type Props = {
    classname?: string;
    link: string;
    classname_text?: string;
    type_button?: "submit" | "reset" | "button";
    onClick?: (e: any) => void;
}
export function Registerbutton({classname,classname_text,link, type_button, onClick}: Props)  {
    return (
        <Button
        id = 'cadastro'
        className={classname}
        asChild
        type={type_button}
        onClick={onClick}
        >
            <Link href={link} className={classname_text}>Cadastro</Link>
        </Button>
    )
}