# recipeAppBackend
# Crate back-end for app on GoLang

Tareas del Back-End (Go)

<!-- Fase 1: Configuración Inicial
    Configura el entorno de desarrollo instalando Go y configurando el GOPATH.
    Inicializa un nuevo módulo Go: go mod init RecipeAppBackend. -->

Fase 2: Conexión a la Base de Datos
    Instala y configura una base de datos (MongoDB o PostgreSQL).
    Usa una biblioteca de ORM como gorm (para PostgreSQL) o mongo-go-driver (para MongoDB) para conectarte a la base de datos.

Fase 3: Modelos de Datos
    Crea estructuras de datos para los modelos:
    User: Contiene Username, Email, Password.
    Recipe: Contiene Title, Ingredients, Instructions, Category, ImageURL.

Fase 4: Rutas y Controladores
    Crea un servidor HTTP usando net/http o un framework como gin.
    Define rutas para gestionar usuarios y recetas (/users, /recipes).
    Implementa controladores para manejar la lógica de negocio.

Fase 5: Autenticación
    Implementa autenticación usando bcrypt para hash de contraseñas.
    Usa jwt-go para crear y verificar tokens de acceso.
    Protege las rutas con un middleware de autenticación.

Fase 6: CRUD de Recetas
    Implementa handlers para crear, leer, actualizar y eliminar recetas.
    Asegura que las operaciones CRUD estén protegidas para usuarios autenticados.

Fase 7: Favoritos
    Modifica la estructura User para incluir favoritos.
    Crea endpoints para añadir y eliminar favoritos.

Fase 8: Filtrado y Búsqueda
    Implementa un endpoint para buscar recetas por título o ingredientes.
    Crea funcionalidad para filtrar recetas por categoría.

Fase 9: Despliegue
    Prepara el back-end para producción.
    Despliega el back-end en una plataforma de hosting como Heroku, Vercel, o DigitalOcean.
    Asegura que todas las funcionalidades estén probadas en producción.
