import { endpoint } from "../endpoint/index.js";
import { withDefaults } from "./with-defaults.js";

export const request = withDefaults(endpoint, {});
