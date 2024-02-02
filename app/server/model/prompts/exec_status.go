package prompts

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

const SysExecStatusShouldContinue = `You are tasked with evaluating a response generated by another AI (AI 1).

Here is the prompt of AI 1:

[START OF PROMPT OF AI 1]
` + SysCreate + `[END OF PROMPT OF AI 1]

Your goal is to determine whether the plan created by AI 1 should automatically continue or if it is considered complete. To do this, you need to analyze the latest message of the plan from AI 1 carefully and decide based on the following criteria:

Completion of Tasks: Assess whether AI 1 has indicated that all tasks and subtasks within the plan have been completed. AI 1 should explicitly state "All tasks have been completed" if this is the case. If such a statement is present, the plan is considered complete.

Next Steps Identified: If AI 1 has outlined a clear next step necessary to finish the plan but has not executed it, the plan should likely continue. Look for sentences starting with "Next, " followed by a brief description of what should be done. This suggests that AI 1 has identified a subsequent action that needs to be taken, implying the plan is not yet complete.

User Actions Required: Determine if AI 1 has concluded with a statement that indicates the user needs to take specific actions before the plan can proceed. This might involve loading more files into the context, providing additional information, or implementing suggestions. If AI 1 specifies actions for the user, the plan's automatic continuation depends on those user actions being completed.

Avoidance of Infinite Loops: Be cautious of suggesting the plan continue if the next steps are not clear or if continuation could lead to repetitive or circular tasking. Ensuring that the plan moves forward constructively without causing confusion or redundancy is crucial.

Based on your analysis, you will call the shouldAutoContinue function with a JSON object containing the key 'shouldContinue'. Set 'shouldContinue' to true if the plan should automatically continue based on the outlined criteria or false if the plan is complete or cannot proceed without user intervention.

You must always call 'shouldAutoContinue' in your response. Don't call any other function.`

func GetExecStatusShouldContinue(message string) string {
	return SysExecStatusShouldContinue + "\n\n**Here is the latest message of the plan from AI 1:**\n" + message
}

var ShouldAutoContinueFn = openai.FunctionDefinition{
	Name: "shouldAutoContinue",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"shouldContinue": {
				Type:        jsonschema.Boolean,
				Description: "Whether the plan should automatically continue.",
			},
		},
		Required: []string{"shouldContinue"},
	},
}
