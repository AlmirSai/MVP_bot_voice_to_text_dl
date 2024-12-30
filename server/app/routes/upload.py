from fastapi import APIRouter, File, UploadFile, HTTPException
from fastapi.responses import HTMLResponse

from server.app.services.json_parser import parse_json
from server.app.services.html_parser import parse_html
from server.app.config import TEMPLATE_DIR


router = APIRouter()


@router.post("/upload-json")
async def upload_json(file: UploadFile = File(...)):
    """
    Endpoint to handle the upload of JSON files.

    This endpoint accepts a JSON file, reads its content, and parses it into a 
    Python dictionary. If the file is not a valid JSON, a 400 HTTP error is 
    raised. For any other exceptions during the file processing, a 500 HTTP 
    error is returned.

    :param file: An UploadFile object representing the uploaded JSON file.
    :return: A dictionary containing a success message and the parsed JSON data.
    :raises HTTPException: If the JSON is invalid or another error occurs.
    """

    try:
        content = await file.read()
        data = parse_json(content)
        return {
            "message": "JSON file processed successfully",
            "data": data
        }
    except ValueError as e:
        raise HTTPException(
            status_code=400,
            detail=str(e)
        )
    except Exception as e:
        raise HTTPException(
            status_code=500,
            detail=str(e)
        )


@router.post("/upload-html")
async def upload_html(file: UploadFile = File(...)):
    """
    Endpoint to handle the upload of HTML files.

    This endpoint accepts an HTML file, reads its content, and extracts all 
    the headers (h1 elements) from it. If there is an error during file 
    processing, a 500 HTTP error is returned.

    :param file: An UploadFile object representing the uploaded HTML file.
    :return: A dictionary containing a success message and the extracted headers.
    :raises HTTPException: If an error occurs during file processing.
    """

    try:
        content = await file.read()
        headers = parse_html(content)
        return {
            "message": "HTML file processed successfully",
            "headers": headers
        }
    except Exception as e:
        raise HTTPException(
            status_code=500,
            detail=str(e)
        )


@router.get("/upload", response_class=HTMLResponse)
async def upload_form():
    """
    Endpoint to serve the upload form HTML page.

    This endpoint returns the HTML content of the upload form page. If the
    HTML template file is not found, a 404 HTTP error is raised.

    :return: An HTMLResponse containing the upload form.
    :raises HTTPException: If the HTML template file is not found.
    """

    try:
        with open(
            TEMPLATE_DIR / "upload_form.html",
            "r",
            encoding="utf-8") as file:
            html_content = file.read()
        return HTMLResponse(content=html_content)
    except FileNotFoundError:
        raise HTTPException(
            status_code=404,
            detail="HTML template not found"
        )
