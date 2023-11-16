import os
import sys
import string
import xml.etree.ElementTree as ET
from openai import OpenAI

# Get command-line arguments for project task name, description, and existing project tasks
project_task_name = sys.argv[1]
project_task_description = sys.argv[2]
existing_projects_tasks = sys.argv[3]

# Initialize the OpenAI client with the API key
client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))

# Create a chat completion request to generate a list of project tasks
completion = client.chat.completions.create(model="gpt-3.5-turbo",
messages=[
    {"role": "system", "content": "You are a team member with a knack for being creative. You are good at coming up with unique and interesting ideas on what tasks your team should work on to accomplish any given project."},
    {"role": "user", "content": "Create a list of 5 specific tasks that will lay out what the team needs to do in order to accomplish a particular project " +
                                f"with this name: {project_task_name} and this description: {project_task_description}. " +
                                f"Also, keep in mind that this information is already contained in the project: {existing_projects_tasks}. " +
                                "Do not add any Arabic numerals (0-9) or punctuation symbols (like '.', ',', ':', etc.) to the task name under any circumstance. " +
                                "Write the tasks in this format: 'Name: Task name, Description: Task description'. Each task must have a name and a description, separated by ', Description: '. " +
                                "Make the task description one or two sentences long. After going through all of the tasks, do not provide any summary or conclusion text. " +
                                "I only want the series of task names and descriptions. " +
                                "Here are some examples: `Name: Determine the scope of the project, Description: Have a team meeting to discuss all of the features that should be included in this project.` " +
                                "`Name: Create a project timeline, Description: Create a timeline that shows when each task should be completed.` " +
                                "`Name: Create a project budget, Description: Create a budget that shows how much money will be spent on each task.` " +
                                "`Name: Create a project schedule, Description: Create a schedule that shows when each task should be completed. Set up short-term and long-term goals for each phase of the project.` "
                                "`Name: Get feedback from the team, Description: Ask the team members for their opinions on what tasks should be included in this project.` " +
                                "`Name: Justify the project, Description: Write a document that explains why this project is important and how it will benefit the company.` " +
                                "`Name: Assign tasks to team members, Description: Assign each task to a team member based on their skills and experience.` "
                                },
])

# Extract the generated task list from the completion response
tasks_str = completion.choices[0].message.content

# Split the task list into individual tasks, separating name and description
tasks = [task.split("Description: ") for task in tasks_str.split("Name: ") if task]

# Create an XML tree for the tasks
root = ET.Element("tasks")

# Iterate through the generated tasks and add them to the XML tree
for task in tasks:
    task_element = ET.SubElement(root, "task")
    name_element = ET.SubElement(task_element, "name")
    
    # Remove digits and punctuation from the task name
    remove_digits_punct = str.maketrans('', '', string.digits + string.punctuation)
    task_name = task[0].replace('\n', '')
    if task_name[-1] == ' ':
        task_name = task_name[:-1]
    name_element.text = task_name.translate(remove_digits_punct)
    
    description_element = ET.SubElement(task_element, "description")
    description_element.text = task[1]

# Print the XML representation of the generated tasks
print(ET.tostring(root, encoding="unicode"))