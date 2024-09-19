import { Input } from "../ui/input";
import { Label } from "../ui/label";
type Props = {
    id: string;
    labeltext: string;
    placeholder: string;
}
export function Textinput({id,labeltext,placeholder}: Props){

    return (
        <div className="flex flex-col space-y-1.5">
            <Label htmlFor={id} className="text-white">{labeltext}</Label>
            <Input inputMode="email" id={id} placeholder={placeholder} className="text-white"/>
          </div>
    )
}