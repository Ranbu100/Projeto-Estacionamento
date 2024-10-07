import * as React from "react";
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from "@/components/ui/select";
import { Controller } from "react-hook-form";

interface DropdownProps {
  label: string;
  options: { label: string, value: number }[];
  control: any;
  name: string;
}

const Dropdown: React.FC<DropdownProps> = ({ label, options, control, name }) => {
  return (
    <div>
      <label className="block text-sm font-medium text-white">{label}</label>
      <Controller
        control={control}
        name={name}
        render={({ field }) => (
          <Select value={field.value} onValueChange={field.onChange}>
            <SelectTrigger className="text-white">
              <SelectValue placeholder="Selecione uma opção" />
            </SelectTrigger>
            <SelectContent>
              {options.map((option, index) => (
                <SelectItem key={index} value={option.value.toString()}>
                  {option.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        )}
      />
    </div>
  );
};

export default Dropdown;
