from fastapi import FastAPI
from routes.upload import router as upload_router

import sys

# Initialize FastAPI app
app = FastAPI(title="FastAPI Project")

# Include routes
app.include_router(upload_router)

@app.get("/")
async def root():
    return {"message": "Welcome to the FastAPI Project!"}

# Dependency check
def check_dependencies():
    try:
        import bs4  # BeautifulSoup dependency
    except ImportError:
        print(
            "Error: Missing dependencies. Install them by running:\n"
            "  pip install -r requirements.txt"
        )
        sys.exit(1)

if __name__ == "__main__":
    # Check dependencies before starting the server
    check_dependencies()

    # Run the server
    import uvicorn
    uvicorn.run("app.main:app", host="127.0.0.1", port=8000, reload=True)
