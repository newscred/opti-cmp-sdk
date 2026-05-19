# opti-cmp

TypeScript SDK for the Optimizely CMP API.

## Install

```bash
npm install opti-cmp
# or
pnpm add opti-cmp
```

## Usage

```typescript
import { OptiCMP } from "opti-cmp";

const client = OptiCMP({
  auth: {
    token: "<auth-token>",
  },
});

// List campaigns
const campaigns = await client.campaigns.listCampaigns({});

// Get a task
const task = await client.tasks.getTask({ id: "task-id" });
```
