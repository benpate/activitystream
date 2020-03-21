function do(input map[interface{}]interface{}) map[string]string {

    result := map[string]string{}

    for key, value := range input {
        
        if key, ok := key.(string); ok {

            if value, ok := value.(string); ok {
                result[key] = value
            }
        }
    }

    return result
}