#!/bin/bash

# Script para añadir comentarios a directorios de clases Go con los temas vistos

# Función principal
main() {
    echo "Bienvenido al asistente de comentarios para directorios de clases Go"
    echo "------------------------------------------------------------"
    
    # Obtener lista de directorios de clases
    class_dirs=($(ls -d Clase* 2>/dev/null | sort -V))
    
    if [ ${#class_dirs[@]} -eq 0 ]; then
        echo "No se encontraron directorios de clases (Clase1, Clase2, etc.)"
        exit 1
    fi
    
    echo "Directorios de clases encontrados:"
    printf '%s\n' "${class_dirs[@]}"
    echo ""
    
    # Procesar cada directorio
    for dir in "${class_dirs[@]}"; do
        process_directory "$dir"
    done
    
    echo "Proceso completado!"
}

# Función para procesar un directorio
process_directory() {
    local dir=$1
    
    echo "Procesando directorio: $dir"
    
    # Verificar si ya tiene comentario
    if [ -f "$dir/.topic" ]; then
        current_comment=$(cat "$dir/.topic")
        echo "Comentario actual: $current_comment"
        read -p "¿Deseas modificarlo? (s/n): " modify
        if [[ "$modify" != "s" ]]; then
            return
        fi
    fi
    
    # Solicitar nuevo comentario
    read -p "Ingresa los temas principales para $dir (ej: 'uso de variables y datos'): " topics
    
    # Guardar comentario en archivo oculto
    echo "$topics" > "$dir/.topic"
    echo "Comentario guardado: $topics"
    echo ""
}

# Ejecutar función principal
main