#!/usr/bin/env python3
import os
from datetime import datetime


def main():
    # Get current date in YYYYMMDD format
    today = datetime.now()
    date_str = today.strftime("%Y%m%d")

    # Create output filename
    output_file = f"unionfind_{date_str}.py"

    # Check if file already exists in parent directory
    output_path = os.path.join("..", output_file)
    if os.path.exists(output_path):
        print(f"File {output_file} already exists for today")
        return

    # Read template file
    template_path = os.path.join("..", "unionfind_template.tmpl")
    try:
        with open(template_path, 'r') as template_file:
            template_content = template_file.read()
    except FileNotFoundError:
        print(f"Error: Template file not found at {template_path}")
        return

    # Create output file in parent directory
    try:
        with open(output_path, 'w') as output_file_handle:
            # Add header comment
            header = f"# Generated on {today.strftime('%Y-%m-%d %H:%M:%S')}\n"
            header += f"# Daily practice file: {output_file}\n\n"
            output_file_handle.write(header)

            # Write template content to output file
            output_file_handle.write(template_content)

        print(f"✅ Successfully generated {output_file} for today's practice!")
    except IOError as e:
        print(f"Error creating output file: {e}")


if __name__ == "__main__":
    main()
