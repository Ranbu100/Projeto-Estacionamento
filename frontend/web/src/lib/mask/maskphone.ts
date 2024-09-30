export const normalizePhoneNumber = (value: string) => {
    // Remove all non-numeric characters
    const cleanedValue = value.replace(/\D/g, '')
    
    // If the resulting string is empty, return an empty string
    if (!cleanedValue) return ''
    
    // Build the formatted string starting with the area code
    let formattedValue = '(' + cleanedValue.substring(0, 2)
    
    // Add the area code, and if there are more than 2 characters, add the first group of 5 digits
    if (cleanedValue.length > 2) {
      formattedValue += ') ' + cleanedValue.substring(2, 7)
    }
    
    // If there are more than 7 characters, add the hyphen and the final group of 4 digits
    if (cleanedValue.length > 7) {
      formattedValue += '-' + cleanedValue.substring(7, 11)
    }
  
    return formattedValue
  }