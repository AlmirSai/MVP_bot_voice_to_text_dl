from fastapi import FastAPI

from routes.upload import router as upload_router

import sys


app = FastAPI(title="FastAPI Project")

app.include_router(upload_router)


@app.get("/")
async def root():
    """
    Welcome to the FastAPI Project!

    This endpoint is the root of the API and just returns a message.
    """
    return {"message": "Welcome to the FastAPI Project!"}


def check_dependencies():
    """
    Checks if the required dependencies are installed. If not, prints an error
    message with instructions on how to install them and exits the program.
    """
    try:
        import bs4
    except ImportError:
        print(
            "Error: Missing dependencies. Install them by running:\n"
            "pip install -r requirements.txt"
        )
        sys.exit(1)


if __name__ == "__main__":
    check_dependencies()

    import uvicorn
    uvicorn.run("app.main:app", host="127.0.0.1", port=8000, reload=True)
