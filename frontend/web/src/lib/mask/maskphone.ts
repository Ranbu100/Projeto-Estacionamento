export const formatPhoneNumber = (value: string): string => {
  if (!value) return value;
  
  // Remove todos os caracteres não numéricos
  const phoneNumber = value.replace(/\D/g, '');

  // Limita o número de dígitos
  if (phoneNumber.length > 11) {
    return phoneNumber.slice(0, 11); // Limita a 11 dígitos numéricos
  }

  // Formatar o número no formato (XX) XXXXX-XXXX
  if (phoneNumber.length <= 10) {
    return phoneNumber.replace(/(\d{2})(\d{4})(\d{0,4})/, '($1) $2-$3');
  } else {
    return phoneNumber.replace(/(\d{2})(\d{5})(\d{0,4})/, '($1) $2-$3');
  }
};


export const unformatPhoneNumber = (value: string): string => {
  return value.replace(/\D/g, '');
};
