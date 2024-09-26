import { FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Senhainput } from "./input_senha";
import { useState } from "react";

export const RenderField = ({ control, name, label, type, errors, ...rest }: any) => {
  const [isVisible, setIsVisible] = useState(false);

  return (
    <FormField control={control} name={name} render={({ field }) => (
      <FormItem>
        <FormLabel className="text-white">{label}</FormLabel>
        <FormControl>
          {type === "senha" || type === "confirmarsenha" ? (
            <Senhainput
              id={type == "senha" ? "senha" : "confirmarsenha"}
              isVisible={isVisible}
              onChangeVisibility={() => setIsVisible(!isVisible)}
              value={field.value || ''}
              onChange={(e) => field.onChange(e.target.value)}
              {...rest}
              
            />
          ) : (
            <Input className="text-white" placeholder={`Escreva seu ${label.toLowerCase()}`} {...field} defaultValue={''} {...rest} />
          )}
        </FormControl>
        <FormMessage>{errors[name]?.message}</FormMessage>
      </FormItem>
    )} />
  );
};
