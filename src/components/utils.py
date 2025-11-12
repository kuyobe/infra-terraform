import os
import logging
import json
import yaml

# Set up logging
logger = logging.getLogger(__name__)

def load_json(file_path):
    """Load JSON file"""
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except FileNotFoundError as e:
        logger.error(f"File not found: {file_path}")
        raise e
    except json.JSONDecodeError as e:
        logger.error(f"Invalid JSON: {file_path}")
        raise e

def load_yaml(file_path):
    """Load YAML file"""
    try:
        with open(file_path, 'r') as file:
            return yaml.safe_load(file)
    except FileNotFoundError as e:
        logger.error(f"File not found: {file_path}")
        raise e
    except yaml.YAMLError as e:
        logger.error(f"Invalid YAML: {file_path}")
        raise e

def create_directory(path):
    """Create directory if it doesn't exist"""
    if not os.path.exists(path):
        os.makedirs(path)