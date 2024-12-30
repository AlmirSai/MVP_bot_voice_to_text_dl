import json


def parse_json(content: bytes):
    """
    Parse a JSON file and return the contents as a Python dictionary.

    Args:
        content (bytes): The contents of the JSON file to be parsed.

    Returns:
        dict: The contents of the JSON file as a Python dictionary.

    Raises:
        ValueError: If the JSON file is invalid.
    """
    try:
        data = json.loads(content.decode("utf-8"))
        return data
    except json.JSONDecodeError:
        raise ValueError("Invalid JSON file")
