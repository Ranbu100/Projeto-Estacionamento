import Link from "next/link";
import { Button } from "../ui/button";

type Props = {
    classname: string;
    link: string;
    classname_text?: string;
}
export function Loginbutton({classname, link, classname_text}: Props){
    return (
        <Button id='logar' variant="outline" className={classname} asChild>
            <Link href={link} className={classname_text}>Login</Link>
        </Button>
    )
}