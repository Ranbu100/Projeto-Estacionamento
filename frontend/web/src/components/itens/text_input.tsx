import { Input } from "../ui/input";
import { Label } from "../ui/label";
type Props = {
    id: string;
    labeltext: string;
    placeholder: string;
    inputmode?: "search" | "text" | "email" | "tel" | "url" | "none" | "numeric" | "decimal";
    value?: string;
    onChange?: (e: any) => void;
}
export function Textinput({id,labeltext,placeholder, inputmode, value, onChange}: Props){

    return (
        <div className="flex flex-col space-y-1.5">
            <Label htmlFor={id} className="text-white">{labeltext}</Label>
            <Input onChange={onChange} value={value} inputMode={inputmode} id={id} placeholder={placeholder} className="text-white"/>
          </div>
    )
}