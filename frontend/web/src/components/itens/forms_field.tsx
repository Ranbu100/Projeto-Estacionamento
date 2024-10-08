'use client'
import { FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Senhainput } from "./input_senha";
import { ChangeEvent, useState } from "react";
import { useForm } from "react-hook-form";


export const RenderField = ({ control, name, label, type, errors, ...rest }: any) => {
  const [isVisible, setIsVisible] = useState(false);
  const { setValue } = useForm();



  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem>
          <FormLabel className="text-white">{label}</FormLabel>
          <FormControl>
            {type === "senha" || type === "confirmarsenha" ? (
              <Senhainput
                id={type === "senha" ? "senha" : "confirmarsenha"}
                isVisible={isVisible}
                onChangeVisibility={() => setIsVisible(!isVisible)}
                value={field.value ?? ''} // Garante que nunca será undefined
                onChange={(e) => field.onChange(e.target.value)}
                {...rest}
              />
            ) : (
              <Input
                type={type == "email" && name != "nome" ? "email" : "tel"}
                className="text-white"
                placeholder={`Escreva seu ${label == "Telefone"? "Telefone(Somente números)".toLowerCase(): label.toLowerCase() }`}
                value={field.value ?? ''}
                onChange={(e) => {
                  field.onChange(e.target.value); // Passa o valor para o react-hook-form
                }}
                {...rest}
              />
            )}
          </FormControl>
          <FormMessage>{errors[name]?.message}</FormMessage>
        </FormItem>
      )}
    />
  );
};
