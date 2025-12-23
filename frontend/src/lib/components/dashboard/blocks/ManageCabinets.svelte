<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import { Building, AlertTriangle, CreditCard } from "@lucide/svelte";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { Admin } from "$lib/types/users/admin";
    import type { SuperAdmin } from "$lib/types/users/superadmin";
    import type { Cabinet } from "$lib/types/cabinet";
    import CabinetSelector from "./cabinets/CabinetSelector.svelte";
    import ManageCabinet from "./ManageCabinet.svelte";

    interface IProps {
        user: SuperAdmin | Admin;
    }

    let { user }: IProps = $props();
    let permissions = getPermissionsFromUserType(user.type);

    let selectedCabinet: Cabinet | null = $state(null);
</script>

<Block Icon={Building} group="manage_cabinet" title="Manage Cabinets">
    {#if selectedCabinet === null}
        <CabinetSelector {user} bind:selectedCabinet />
    {:else}
        <ManageCabinet
            {user}
            cabinet={selectedCabinet}
            onBack={() => {
                selectedCabinet = null;
            }}
        />
    {/if}
</Block>
