FROM gitpod/workspace-full:latest

USER root
# Install custom tools, runtime, etc.
RUN apt-get update && apt-get install -y \
    && apt-get install git-lfs -y \
    && apt-get install --no-install-recommends \
        # apt-utils \
        ca-certificates \
        libgl1-mesa-dev \
        xorg-dev \
        gcc -y \
    && apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*


USER gitpod
# Apply user-specific settings
ENV GO111MODULE on

# Give back control
USER root
