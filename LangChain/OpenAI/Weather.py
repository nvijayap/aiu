# https://docs.langchain.com/oss/python/langchain/overview

# pip install -qU langchain "langchain[openai]"

import json
import sys
import json_repair
from langchain.agents import create_agent

def get_weather(city: str) -> str:
    """Get weather for a given city."""
    return f"It's always sunny in {city}!"

agent = create_agent(
    model="openai:gpt-5.4",
    tools=[get_weather],
    system_prompt="You are a helpful assistant",
)

result = agent.invoke(
    {"messages": [{"role": "user", "content": "What's the weather in San Francisco?"}]}
)

fake_json = str(result["messages"][-1].content_blocks)

real_json = json_repair.repair_json(fake_json)

print(real_json)
