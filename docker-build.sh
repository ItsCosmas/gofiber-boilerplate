build (){
    # Package the backend source to a build folder
    rsync -av --progress . build \
    --exclude mongodata \
    --exclude postgres_data \
    --exclude .git \
    --exclude src/tmp \
    --exclude 'src/.env.example' \
    --exclude 'README.md' \
    --exclude 'Dockerfile' \
    --exclude '.air.toml' \
    --exclude '.dockerignore' \
    --exclude 'docker-compose.yml' \
    --exclude 'LICENSE' \
    --exclude '.gitignore'

    # Move to build directory
    cd build
    # Docker Build
    # Specify docker file
    # Tag the image with a user provided version, useful when pushing the image and to identify it
    docker build -f Dockerfile.prod . -t $1
    # Move Back to root
    cd ..
    # Delete the build directory
    rm -rf build
}

# Allow User to pass version information to script
while getopts v: flag
do
    case "${flag}" in
        v) version=${OPTARG};;
    esac
done

build $version;