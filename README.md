# Promote an image tag with PR:s

Different ideas for how to update the image tags on different deployments.
Assume that we have two environments: dev and prod. Goal: The image is updated
automatically in the gitops repository for dev when a new image is available. A
PR is created to make the same change to the gitops repo for prod.

## Examples of stuff that can trigger the promotion

- The pipeline that pushes the image, e.g. on every commit or when a git tag is
  pushed
- Webhook from image registry when a new image is published

```mermaid
flowchart LR
    A1([Update to application triggers image build]) -.-> B([Bot promotes dev])
    A2([New git tag triggers image build]) -.-> B
    A3([Image registry notifies that a new image exists]) -.-> B
    B --> C([Bot creates PR to promote prod])
    C --> D([Non-bot merges PR])
```
