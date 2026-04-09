#!/usr/bin/env python3
import os
from datetime import datetime


def main():
    # Get current date in YYYYMMDD format
    today = datetime.now()
    date_str = today.strftime("%Y%m%d")

    # Define template configurations
    templates = {
        "traverse": {"template_file": "traverse_template.tmpl", "output_prefix": "traverse"},
        "dag": {"template_file": "dag_template.tmpl", "output_prefix": "dag"},
        "dijkstra": {"template_file": "dijkstra_template.tmpl", "output_prefix": "dijkstra"},
    }

    generated_count = 0
    generated_files = []

    # Generate all templates
    for template_type, config in templates.items():
        # Create output filename
        output_file = f"{config['output_prefix']}_{date_str}.py"

        # Check if file already exists in parent directory
        output_path = os.path.join("..", output_file)
        if os.path.exists(output_path):
            print(f"File {output_file} already exists for today")
            continue

        # Read template file
        template_path = os.path.join("..", config['template_file'])
        try:
            with open(template_path, 'r') as template_file:
                template_content = template_file.read()
        except FileNotFoundError:
            print(f"⚠️  Warning: Could not open template file {config['template_file']}")
            continue

        # Create output file in parent directory
        try:
            with open(output_path, 'w') as output_file_handle:
                # Add header comment
                header = f"# Generated on {today.strftime('%Y-%m-%d %H:%M:%S')}\n"
                header += f"# Daily practice file: {output_file} ({template_type} template)\n\n"
                output_file_handle.write(header)

                # Write template content to output file
                output_file_handle.write(template_content)

            generated_files.append(output_file)
            generated_count += 1
        except IOError as e:
            print(f"⚠️  Warning: Could not create output file {output_file}: {e}")
            continue

    # Print summary
    if generated_count == 0:
        print("No new files generated - all templates already exist for today")
    else:
        print(f"✅ Successfully generated {generated_count} file(s) for today's practice:")
        for file in generated_files:
            print(f"   - {file}")


if __name__ == "__main__":
    main()
