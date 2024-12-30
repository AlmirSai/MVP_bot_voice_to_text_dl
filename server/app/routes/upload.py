from fastapi import APIRouter, File, UploadFile, HTTPException
from fastapi.responses import HTMLResponse

from server.app.services.json_parser import parse_json
from server.app.services.html_parser import parse_html
from server.app.config import TEMPLATE_DIR


router = APIRouter()


@router.post("/upload-json")
async def upload_json(file: UploadFile = File(...)):
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
    try:
        with open(TEMPLATE_DIR / "upload_form.html", "r", encoding="utf-8") as file:
            html_content = file.read()
        return HTMLResponse(content=html_content)
    except FileNotFoundError:
        raise HTTPException(
            status_code=404,
            detail="HTML template not found"
        )