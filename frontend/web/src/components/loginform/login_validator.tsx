import * as yup from "yup";

export const formSchema = yup.object().shape({
  email: yup.string().email("Email inválido").required("Email é obrigatório"),
  senha: yup.string().min(6, "A senha deve ter pelo menos 6 caracteres").required("Senha é obrigatória"),
});
