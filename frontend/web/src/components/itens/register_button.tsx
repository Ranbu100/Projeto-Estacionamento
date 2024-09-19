import Link from "next/link";
import { Button } from "../ui/button"

type Props = {
    classname?: string;
    link: string;
    classname_text?: string;
}
export function Registerbutton({classname,classname_text,link}: Props)  {
    return (
        <Button
        id = 'cadastro'
        className={classname}
        asChild
        >
            <Link href={link} className={classname_text}>Cadastro</Link>
        </Button>
    )
}