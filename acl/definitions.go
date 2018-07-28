/*
 * Copyright (c) 2013-2018, Jeremy Bingham (<jeremy@goiardi.gl>)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package acl

import (

)

// Define the casbin RBAC model and the skeletal $$default$$ policy.

const modelDefinition = `[request_definition]
r = sub, obj, kind, subkind, act

[policy_definition]
p = sub, obj, kind, subkind, act, eft

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = r.sub == "pivotal" && p.eft != "deny" || g(r.sub, p.sub) && r.kind == p.kind && r.subkind == p.subkind && r.obj == p.obj && r.act == p.act
`

// NOTE: MySQL/Postgres implementations of this may require some mild heroics
// to put convert this to a form suitable to put in the DB. We'll see what ends
// up happening.

// group, subkind, kind, name, perm, effect

const defaultPolicySkel = `p, role##admins, containers, containers, $$default$$, create, allow
p, role##admins, containers, containers, $$default$$, read, allow
p, role##users, containers, containers, $$default$$, read, allow
p, role##admins, containers, containers, $$default$$, update, allow
p, role##admins, containers, containers, $$default$$, delete, allow
p, role##admins, containers, containers, $$default$$, grant, allow
p, role##users, containers, containers, clients, delete, allow
p, role##users, containers, containers, nodes, create, allow
p, role##users, containers, containers, environments, create, allow

p, role##admins, $$root$$, containers, $$default$$, create, allow
p, role##admins, $$root$$, containers, $$default$$, read, allow
p, role##users, $$root$$, containers, $$default$$, read, allow
p, role##admins, $$root$$, containers, $$default$$, update, allow
p, role##admins, $$root$$, containers, $$default$$, delete, allow
p, role##admins, $$root$$, containers, $$default$$, grant, allow

p, role##admins, groups, containers, $$default$$, create, allow
p, role##admins, groups, containers, $$default$$, read, allow
p, role##admins, groups, containers, $$default$$, update, allow
p, role##admins, groups, containers, $$default$$, delete, allow
p, role##admins, groups, containers, $$default$$, grant, allow
p, role##users, groups, containers, clients, read, deny

p, role##admins, cookbooks, containers, $$default$$, create, allow
p, role##admins, cookbooks, containers, $$default$$, read, allow
p, role##admins, cookbooks, containers, $$default$$, update, allow
p, role##admins, cookbooks, containers, $$default$$, delete, allow
p, role##admins, cookbooks, containers, $$default$$, grant, allow
p, role##users, cookbooks, containers, $$default$$, create, allow
p, role##users, cookbooks, containers, $$default$$, read, allow
p, role##users, cookbooks, containers, $$default$$, update, allow
p, role##users, cookbooks, containers, $$default$$, delete, allow
p, role##clients, cookbooks, containers, $$default$$, read, allow

p, role##admins, environments, containers, $$default$$, create, allow
p, role##admins, environments, containers, $$default$$, read, allow
p, role##admins, environments, containers, $$default$$, update, allow
p, role##admins, environments, containers, $$default$$, delete, allow
p, role##admins, environments, containers, $$default$$, grant, allow
p, role##users, environments, containers, $$default$$, create, allow
p, role##users, environments, containers, $$default$$, read, allow
p, role##users, environments, containers, $$default$$, update, allow
p, role##users, environments, containers, $$default$$, delete, allow
p, role##clients, environments, containers, $$default$$, read, allow

p, role##admins, roles, containers, $$default$$, create, allow
p, role##admins, roles, containers, $$default$$, read, allow
p, role##admins, roles, containers, $$default$$, update, allow
p, role##admins, roles, containers, $$default$$, delete, allow
p, role##admins, roles, containers, $$default$$, grant, allow
p, role##users, roles, containers, $$default$$, create, allow
p, role##users, roles, containers, $$default$$, read, allow
p, role##users, roles, containers, $$default$$, update, allow
p, role##users, roles, containers, $$default$$, delete, allow
p, role##clients, roles, containers, $$default$$, read, allow

p, role##admins, data, containers, $$default$$, create, allow
p, role##admins, data, containers, $$default$$, read, allow
p, role##admins, data, containers, $$default$$, update, allow
p, role##admins, data, containers, $$default$$, delete, allow
p, role##admins, data, containers, $$default$$, grant, allow
p, role##users, data, containers, $$default$$, create, allow
p, role##users, data, containers, $$default$$, read, allow
p, role##users, data, containers, $$default$$, update, allow
p, role##users, data, containers, $$default$$, delete, allow
p, role##clients, data, containers, $$default$$, read, allow

p, role##admins, nodes, containers, $$default$$, create, allow
p, role##admins, nodes, containers, $$default$$, read, allow
p, role##admins, nodes, containers, $$default$$, update, allow
p, role##admins, nodes, containers, $$default$$, delete, allow
p, role##admins, nodes, containers, $$default$$, grant, allow
p, role##users, nodes, containers, $$default$$, create, allow
p, role##users, nodes, containers, $$default$$, read, allow
p, role##users, nodes, containers, $$default$$, update, allow
p, role##users, nodes, containers, $$default$$, delete, allow
p, role##clients, nodes, containers, $$default$$, create, allow
p, role##clients, nodes, containers, $$default$$, read, allow

p, role##admins, clients, containers, $$default$$, create, allow
p, role##admins, clients, containers, $$default$$, read, allow
p, role##admins, clients, containers, $$default$$, update, allow
p, role##admins, clients, containers, $$default$$, delete, allow
p, role##admins, clients, containers, $$default$$, grant, allow
p, role##users, clients, containers, $$default$$, read, allow
p, role##users, clients, containers, $$default$$, delete, allow

p, role##admins, sandboxes, containers, $$default$$, create, allow
p, role##admins, sandboxes, containers, $$default$$, read, allow
p, role##admins, sandboxes, containers, $$default$$, update, allow
p, role##admins, sandboxes, containers, $$default$$, delete, allow
p, role##admins, sandboxes, containers, $$default$$, grant, allow
p, role##users, sandboxes, containers, $$default$$, create, allow

p, role##admins, log-infos, containers, $$default$$, create, allow
p, role##admins, log-infos, containers, $$default$$, read, allow
p, role##admins, log-infos, containers, $$default$$, update, allow
p, role##admins, log-infos, containers, $$default$$, delete, allow
p, role##admins, log-infos, containers, $$default$$, grant, allow
p, role##users, log-infos, containers, $$default$$, create, allow

p, role##admins, reports, containers, $$default$$, create, allow
p, role##admins, reports, containers, $$default$$, read, allow
p, role##admins, reports, containers, $$default$$, update, allow
p, role##admins, reports, containers, $$default$$, delete, allow
p, role##admins, reports, containers, $$default$$, grant, allow
p, role##users, reports, containers, $$default$$, create, allow
p, role##clients, reports, containers, $$default$$, create, allow

p, role##admins, shoveys, containers, $$default$$, create, allow
p, role##admins, shoveys, containers, $$default$$, read, allow
p, role##admins, shoveys, containers, $$default$$, update, allow
p, role##admins, shoveys, containers, $$default$$, delete, allow
p, role##admins, shoveys, containers, $$default$$, grant, allow
p, role##clients, shoveys, containers, $$default$$, update, allow

p, role##billing-admins, billing-admins, groups, $$default$$, read, allow
p, role##billing-admins, billing-admins, groups, $$default$$, update, allow

p, role##admins, admins, groups, $$default$$, create, allow
p, role##admins, admins, groups, $$default$$, read, allow
p, role##admins, admins, groups, $$default$$, update, allow
p, role##admins, admins, groups, $$default$$, delete, allow
p, role##admins, admins, groups, $$default$$, grant, allow

p, role##admins, clients, groups, $$default$$, create, allow
p, role##admins, clients, groups, $$default$$, read, allow
p, role##admins, clients, groups, $$default$$, update, allow
p, role##admins, clients, groups, $$default$$, delete, allow
p, role##admins, clients, groups, $$default$$, grant, allow

p, role##admins, users, groups, $$default$$, create, allow
p, role##admins, users, groups, $$default$$, read, allow
p, role##admins, users, groups, $$default$$, update, allow
p, role##admins, users, groups, $$default$$, delete, allow
p, role##admins, users, groups, $$default$$, grant, allow

p, role##admins, $$default$$, groups, $$default$$, create, allow
p, role##admins, $$default$$, groups, $$default$$, read, allow
p, role##admins, $$default$$, groups, $$default$$, update, allow
p, role##admins, $$default$$, groups, $$default$$, delete, allow
p, role##admins, $$default$$, groups, $$default$$, grant, allow
p, role##users, $$default$$, groups, $$default$$, read, allow
`