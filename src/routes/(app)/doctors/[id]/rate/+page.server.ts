import type { PageServerLoad } from './$types';

/**
 * Loads the page data based on the provided parameters.
 */
export const load: PageServerLoad = async ({ params }) => {
    return {
        id: params.id
    };
};
